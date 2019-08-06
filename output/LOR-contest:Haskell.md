Вся программа в одном файле.

    module Main where

    import Data.Map hiding (map, filter, delete, null)
    import Data.List
    import System.IO

    data Point = Point {
                        x:: Int,
                        y:: Int }
                 deriving (Eq, Ord)

    toUp (Point x y) = Point x (y + 1)
    toDown (Point x y) = Point x (y - 1)
    toLeft (Point x y) = Point (x + 1) y
    toRight (Point x y) = Point (x - 1) y

    data World = World
        { sizex:: Int,
          sizey:: Int,
          walls:: [Point],
          items:: [Point],
          player:: Point,
          carry:: Bool
        }

    readWorld y [] = World 0 (y - 1) [] [] (Point 0 0) False
    readWorld y (line:rest) = readLine (readWorld (y + 1) rest) 1 line
        where
          readLine world x [] = world { sizex = max x (sizex world) }
          readLine world x (cell:rest) =
              readLine readCell (x + 1) rest
                  where
                    readCell =
                        world {
                               walls = if cell == '#' then point:worldWalls else worldWalls,
                               items = if cell == '.' then point:worldItems else worldItems,
                               player = if cell == '@' then point else (player world) }
                            where
                              point = Point x y
                              worldWalls = walls world
                              worldItems = items world

    getWorld = do
      worldStream <- readFile "data"
      return (readWorld 1 $ lines worldStream)

    showWorld _ 0 = []
    showWorld world y = (showLine (Point (sizex world) y)) : showWorld world (y - 1)
        where
          showLine (Point 0 _) = []
          showLine point = showCell : showLine (toRight point)
              where
                showCell
                    | point `elem` (walls world) = '#'
                    | point == (player world) = '@'
                    | point `elem` (items world) = '.'
                    | otherwise = ' '

    isValid world point@(Point x y) =
        point `notElem` (walls world) &&
        x>0 && x<=(sizex world) &&
        y>0 && y<=(sizey world)

    bringAll world =
        foldr (++) [] (map bringItem $ items world)
        where
          start = player world
          neighbours point = [dir point |
                              dir <- [toUp, toDown, toLeft, toRight],
                              isValid world (dir point) ]
          waveMap [] wave = wave
          waveMap front wave =
              let source = head front
                  nb = filter (`notMember` wave) $ neighbours source
                  g = wave ! source + 1 in
              waveMap (tail front ++ nb)
                          (foldr
                           (\point wm -> insertWith (_ n -> n) point g wm)
                           wave nb)
          waveMap' = waveMap [start] (singleton start 0)
          bringItem item =
              (traceProgram $ reverse itemTrace) ++ ["take"] ++ traceProgram itemTrace ++ ["drop"]
              where
                itemTrace = trace item
                trace item
                   | start == item = [item]
                   | otherwise = item : trace nearest
                   where
                     nearest = minimumBy genOrder $ neighbours item
                         where
                           findGen p = findWithDefault (waveMap' ! item) p waveMap'
                           genOrder p1 p2 = compare (findGen p1) (findGen p2)
                traceProgram [point] = []
                traceProgram (point:rest) = (command point (head rest)):(traceProgram rest)
                    where
                      command (Point x1 y1) (Point x2 y2) =
                          case Point (x2-x1) (y2-y1) of
                            Point 0 1 -> "up"
                            Point 0 (-1) -> "down"
                            Point 1 0 -> "left"
                            Point (-1) 0 -> "right"

    getProgram = do
      str <- getLine
      return (words str)

    applyProgram program world initial =
        foldl applyCommand world program
        where
          applyCommand world str =
              case str of
                "up" -> step toUp
                "down" -> step toDown
                "left" -> step toLeft
                "right" -> step toRight
                "take" -> grab True delete
                "drop" -> grab False (:)
                "reset" -> initial
                "collect" -> applyProgram (bringAll world) world initial
                otherwise -> world
              where
                current = player world
                step dir
                    | isValid world target = world { player = target }
                    | otherwise = world
                    where
                      target = dir current
                grab carryFlag action
                    | carry world == carryFlag = world
                    | otherwise = world { carry=carryFlag,
                                          items=action current (items world) }

    doRepl world initial = do
      putStr (unlines $ showWorld world (sizey world))
      program <- getProgram
      if null program
        then
          return ()
        else
          doRepl (applyProgram program world initial) initial

    main = do
      world <- getWorld
      doRepl world world

[Category:LOR-contest](Category:LOR-contest "wikilink")