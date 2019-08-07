Программа состоит из четырех файлов: adv_world.ml, advserver.ml,
hclient.ml, roboclient.ml. Кроме этого, еще два файла, которые не вошли
в зачет -- adv_world.mli и Makefile.

### \[adv_world.ml\]

    let find map c =
      let rec find_aux i =
        if i = Array.length map then raise Not_found
        else try (i, String.index map.(i) c) with Not_found -> find_aux (i + 1)
      in find_aux 0

    let get_items map =
      let rec get_items_aux accu = function
          (i, j) when i = Array.length map      -> accu
        | (i, j) when j = String.length map.(i) -> get_items_aux accu (i + 1, 0)
        | (i, j) -> get_items_aux (if map.(i).[j] = '$' then (i, j) :: accu else accu) (i, j + 1)
      in
      get_items_aux [] (0, 0) ;;

    let num_of_items map = List.length (get_items map) ;;

### \[advserver.ml\]

    let serve_cmd map ((x, y), is_empty) = function
        "pickup" when is_empty && map.(x).[y] = '$' ->
          map.(x).[y] <- ' ' ;
          ((x, y), false)
      | "drop" when not is_empty && map.(x).[y] <> '$' ->
          if map.(x).[y] <> '.' then map.(x).[y] <- '$' ;
          ((x, y), true)
      | cmd ->
          let cmds = [("north", (-1, 0)); ("south", (1, 0)); ("west", (0, -1)); ("east", (0, 1))] in
          let x', y' = (function (dx, dy) -> x + dx, y + dy) (try List.assoc cmd cmds with Not_found -> 0, 0) in
          ( if (try ignore map.(x).[y]; false with Invalid_argument _ -> true) || map.(x').[y'] = '#' then
              x, y else x', y' ), is_empty

    let send_stat out_channel map ((x, y), is_empty) =
      let map' = Array.map (function line -> String.copy line) map in
      map'.(x).[y] <- '@' ;
      Marshal.to_channel out_channel map' [] ;
      Marshal.to_channel out_channel
        [ if Adv_world.num_of_items map = 0 && is_empty then "You won!" else "Game is in progress";
          if is_empty then "Your hand are empty" else "You carry an item";
          if map.(x).[y] = '$' then "Here lies 1 item" else "Here lies 0 items" ] [] ;
      flush out_channel

    let rec input ic =
      try let line = input_line ic in line :: input ic with End_of_file -> []

    let serve in_channel out_channel =
      let map = Array.of_list (input (open_in Sys.argv.(1))) in
      let rec loop state =
        send_stat out_channel map state ;
        try loop (serve_cmd map state (input_line in_channel)) with
          End_of_file -> ()
      in
      loop ((Adv_world.find map '.'), true) ;;

    Unix.establish_server serve (Unix.ADDR_INET (Unix.inet_addr_loopback, int_of_string Sys.argv.(2))) ;;

### \[hclient.ml\]

    let in_channel, out_channel = Unix.open_connection
      (Unix.ADDR_INET ((Unix.inet_addr_of_string Sys.argv.(1)), int_of_string Sys.argv.(2)))
    in
    while true do
      Array.iter (Printf.printf "%s\n") (Marshal.from_channel in_channel : string array) ;
      List.iter (Printf.printf "%s\n") (Marshal.from_channel in_channel : string list) ;
      flush stdout ;
      Printf.fprintf out_channel "%s\n" (read_line ()) ;
      flush out_channel
    done ;;

\[roboclient.ml\]

    let dfs map (sx, sy) =
      let dxy = [(-1, 0); (1, 0); (0, -1); (0, 1)] in
      let f = Array.init
        (Array.length map) (function i -> Array.make (String.length map.(i)) (-1))
      in
      let q = Queue.create () in
      Queue.push (sx, sy) q ; f.(sx).(sy) <- 0 ;
      while not (Queue.is_empty q) do
        let x, y = Queue.pop q in
        List.iter
          (function (x', y') -> Queue.push (x', y') q ; f.(x').(y') <- f.(x).(y) + 1)
          ( List.filter
              ( function (x', y') ->
                  (try ignore map.(x').[y']; true with Invalid_argument _ -> false) &&
                    map.(x').[y'] <> '#' && f.(x').(y') = -1 )
              (List.map (function (dx, dy) -> x + dx, y + dy) dxy) )
      done ;
      f

    let find_path f (x, y) =
      let rec find_path' (x, y) path rev_path =
        let dxy = [ ((-1, 0), ("south", "north")); ((1, 0), ("north", "south"));
                    ((0, -1), ("east", "west")); ((0, 1), ("west", "east")) ] in
        if f.(x).(y) = 0 then path @ ["pickup"] @ (List.rev rev_path) @ ["drop"] else
          let find_path_aux (dx, dy) (dir, rev_dir) =
            if try f.(x + dx).(y + dy) + 1 = f.(x).(y) with Invalid_argument _ -> false then
              find_path' (x + dx, y + dy) (dir :: path) (rev_dir :: rev_path)
            else []
          in
          List.fold_right
            (function (d, dirs) -> (function [] -> find_path_aux d dirs | path -> path)) dxy []
      in
      find_path' (x, y) [] [] ;;

    let in_channel, out_channel = Unix.open_connection
      (Unix.ADDR_INET ((Unix.inet_addr_of_string Sys.argv.(1)), int_of_string Sys.argv.(2)))
    in

    let map = (Marshal.from_channel in_channel : string array) in
    Array.iter (Printf.printf "%s\n") map ;
    List.iter (Printf.printf "%s\n") (Marshal.from_channel in_channel : string list) ;

    List.iter
      ( function (x, y) ->
          List.iter
            ( function cmd ->
                Printf.fprintf out_channel "%s\n" cmd ;
                flush out_channel ;
                Array.iter (Printf.printf "%s\n") (Marshal.from_channel in_channel : string array) ;
                List.iter (Printf.printf "%s\n") (Marshal.from_channel in_channel : string list) )
            (find_path (dfs map (Adv_world.find map '@')) (x, y)) )
      ( Adv_world.get_items map ) ;;

### \[adv_world.mli\]

    val find : string array -> char -> int * int
    val get_items : string array -> (int * int) list
    val num_of_items : string array -> int

### \[Makefile\]

    adventure: advserver hclient roboclient

    advserver: adv_world.cmo advserver.ml
        ocamlc unix.cma adv_world.cmo advserver.ml -o advserver
    hclient: adv_world.cmo hclient.ml
        ocamlc -g unix.cma adv_world.cmo hclient.ml -o hclient
    roboclient: adv_world.cmo roboclient.ml
        ocamlc -g unix.cma adv_world.cmo roboclient.ml -o roboclient
    adv_world.cmo: adv_world.mli adv_world.ml
        ocamlc -g -c adv_world.mli
        ocamlc -g -c adv_world.ml

    clean:
        rm -Rf *~ *.cm[io] advserver hclient roboclient

[Category:LOR-contest](Category:LOR-contest)