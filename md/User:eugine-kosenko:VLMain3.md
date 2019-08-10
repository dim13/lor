    cons 0 = Nil
    cons n = Cons n (cons (n-1))
    main3 = main3p Nil (cons n)
       where
          main3p 0 = ScalarProduct
          main3p i a = main3p (i-1) (Cons i a)
