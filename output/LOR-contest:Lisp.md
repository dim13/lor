В отличие от других решений, решение на Lisp удалось сделать в одном
файле. После загрузки пакета достаточно вызвать функцию advserver
для сервера и hclient для клиента. Программа для робота запускается,
как команда клиента collect.

    (require 'trivial-sockets)

    ;; World definition and inspection

    (defmacro def-world (&rest attr-names)
      `(progn
          ,@(mapcar #'(lambda (attr-name) `(defparameter ,attr-name nil)) attr-names)
          (defun clear-world ()
             (setf ,@(mapcan #'(lambda (attr-name) (list attr-name nil)) attr-names)))
          (defun pack-world ()
             (list 'setf
                     ,@(mapcan
                         #'(lambda (attr-name)
                              (list `(quote ,attr-name) `(list 'quote ,attr-name))) attr-names)))))

    (def-world
      bound-x bound-y wall-points thing-points entry-point hero-point hero-carries-thing-p)

    (defparameter world-image-name "1")

    (defconstant +directions+
        '((#c(1 0) . go-east)
          (#c(-1 0) . go-west)
          (#c(0 -1) . go-north)
          (#c(0 1) . go-south)))

    (defun show-world ()
      (loop
        for y upto bound-y
        do (loop
             for x upto bound-x
             for point = (complex x y)
             do (princ
                  (cond
                     ((member point wall-points) #\#)
                     ((eql point hero-point) #\@)
                     ((member point thing-points) #\$)
                     (t " ")))
             finally (terpri)))
      (format t "~A~%~A~%~D item(s) here~%"
                 (if (or hero-carries-thing-p
                            (remove entry-point thing-points))
                      "Game is in progress"
                      "You won!")
                 (if hero-carries-thing-p
                      "You carry an item"
                      "Your hands are empty")
                 (count hero-point thing-points)))

    ;; New Path finder

    (defun make-program (directions)
      (mapcar
        #'(lambda (direction)
             `(on-server
                (quote (,(if (numberp direction)
                                 (cdr (assoc direction +directions+))
                                 direction)))))
        `(,@(reverse directions)
          pickup-thing
          ,@(mapcar #'- directions)
          drop-thing)))

    (defun collect ()
      (loop
        with path-table = (make-hash-table)
        for front = () then (cdr front)
        and point = hero-point then (car front)
        while point
        do (loop
             for direction in (mapcar #'car +directions+)
             for neighbour = (+ point direction)
             unless (or
                        (< (realpart neighbour) 0)
                        (> (realpart neighbour) bound-x)
                        (< (imagpart neighbour) 0)
                        (> (imagpart neighbour) bound-y)
                        (member neighbour wall-points)
                        (gethash neighbour path-table))
             do (setf (gethash neighbour path-table) (cons direction (gethash point path-table)))
             (push neighbour front))
        finally
        (return
          `(progn
             ,@(mapcan
                  #'(lambda (point)
                        (make-program (gethash point path-table)))
                  thing-points)))))

    ;; Command interpreter

    (defun reload-world ()
      (clear-world)
      (with-open-file (file world-image-name)
        (loop
          initially (setq bound-x 0)
          for line = (read-line file nil nil)
          and y upfrom 0
          while line
          do (loop
               for c across line
                and x upfrom 0
               do (case c
                   (#\# (push (complex x y) wall-points))
                   (#\$ (push (complex x y) thing-points))
                   (#\. (setf hero-point (setf entry-point (complex x y)))))
                finally (setq bound-x (max bound-x x)))
          finally (setq bound-y y))))

    (defun load-world (file-name)
      (setf world-image-name file-name)
      (reload-world))

    ;; Go commands
    (defmacro def-commands (directions)
      `(progn
         ,@(mapcar #'(lambda (pair)
                            `(defun ,(cdr pair) ()
                              (let ((new-point (+ hero-point ,(car pair))))
                                 (unless (member new-point wall-points)
                                    (setf hero-point new-point)))))
                      directions)))

    (def-commands #.+directions+)

    (defun pickup-thing ()
      (when
        (and (member hero-point thing-points)
              (not hero-carries-thing-p))
        (setf hero-carries-thing-p t
                thing-points (remove hero-point thing-points :count 1))))

    (defun drop-thing ()
      (when hero-carries-thing-p
         (setf hero-carries-thing-p nil)
         (push hero-point thing-points)))

    ;; Network layer

    (defun send (stream command)
         (write-line (format nil "~S~%" command) stream)
         (force-output stream))

    ;; Server

    (defun advserver (file-name port)
      (load-world file-name)
      (trivial-sockets:with-server
        (server (:port port :reuse-address t))
        (loop
         (with-open-stream
          (stream (trivial-sockets:accept-connection server))
          (eval (read stream))
          (send stream (pack-world))))))

    ;; Client

    (defparameter host "localhost")
    (defparameter port 7766)

    (defun on-server (server-program)
      (with-open-stream
        (stream (trivial-sockets:open-stream host port))
        (send stream server-program)
        (eval (read stream))
        (show-world)))

    (defun hclient (l-host l-port)
      (setf host l-host port l-port)
      (on-server '(reload-world))
      (loop
        (princ "> ") (force-output)
        (eval (eval (read)))))

[Category:LOR-contest](Category:LOR-contest "wikilink")
[Category:Lisp](Category:Lisp "wikilink")