Так же, как и для [Ocaml](LOR-contest:Ocaml "wikilink"), программа
состоит из четырех файлов -- advworld.py, advserver.py,
hclient.py, roboclient.py.

### \[advworld.py\]

    from copy import deepcopy
    
    class World:
        def __init__(self, world_file):
            self.map = []
            self.items = []
            
            f = open(world_file)
            
            row = 0
            while True:
                l = f.readline().rstrip()
                
                if l == '':
                    break
                
                map_line = str(l).replace('$', ' ').replace('.', ' ')
                self.map.append(list(map_line))
                
                for col in xrange(len(l)):
                    if l[col] == '$':
                        self.items.append((row, col))
                    if l[col] == '.':
                        self.entry = (row, col)
                
                row += 1
            
            self.height = len(self.map)
            self.width = max(len(r) for r in self.map)
            self.hands_empty = True
            self.position = self.entry
            
        def gameWon(self):
            ok = True
            for item in self.items:
                if item != self.entry:
                    ok = False
                    break
            return ok and self.hands_empty
        
        def move(self, dx, dy):
            new_col = self.position[1] + dx
            new_row = self.position[0] + dy
            
            try:
                new_place = self.map[new_row][new_col]
                if new_place == ' ':
                    self.position = (new_row, new_col)
            except IndexError:
                pass
                
        def action_north(self):
            self.move(0, -1)
        
        def action_south(self):
            self.move(0, 1)
        
        def action_east(self):
            self.move(1, 0)
        
        def action_west(self):
            self.move(-1, 0)
        
        def action_pickup(self):
            if self.position in self.items:
                self.items.remove(self.position)
            self.hands_empty = False
        
        def action_drop(self):
            if not self.hands_empty:
                self.hands_empty = True
                self.items.append(self.position)
        
        def performActions(self, actions):
            states = []
            for action in actions:
                try:
                    getattr(self, 'action_'+action)()
                    states.append(deepcopy(self))
                except KeyError:
                    pass
            return states
        
        def __str__(self):
            s = []
            
            for row in xrange(self.height):
                line = ''
                for col in xrange(self.width):        
                    pos = (row, col)
                    try:
                        char = self.map[row][col]
                    
                        if pos == self.position:
                            char = '@'
                        elif pos == self.entry:
                            char = '.'
                        elif pos in self.items:
                            char = '$'
                        line += char
                    except IndexError:
                        pass
                s.append(line)
                
            if self.gameWon():
                s.append('You won!')
            else:
                s.append('Game is in progress')
            
            if self.hands_empty:
                s.append('Your hands are empty')
            else:
                s.append('You carry an item')
            
            s.append('Here lie(s) %d item(s)' % (self.items.count(self.position)))
            
            return '\n'.join(s)

### \[advserver.py\]

    from sys import argv
    from  Pyro.core import Daemon, initServer, ObjBase
    from advworld import World
    
    class Server(ObjBase):
        def init(self, fname):
            self.world = World(fname)
        
        def read_world(self):
            return self.world
        
        def perform_actions(self, actions):
            return self.world.performActions(actions)
    
    if __name__ == '__main__':
        server = Server()
        server.init(argv[1])
        
        initServer()
        daemon = Daemon(port=int(argv[2]))
        daemon.connect(server, 'adventure_server')
        daemon.requestLoop()

### \[hclient.py\]

    from sys import stdin, stdout, argv
    import Pyro.core
    from roboclient import Robot
    
    if __name__ == '__main__':
        server = Pyro.core.getProxyForURI(
            'PYROLOC://%s:%s/adventure_server' % (argv[1], argv[2]))
        
        print server.read_world()    
        while True:
            stdout.write('> ')
            commands = stdin.readline().strip().lower().split()
            
            for cmd in commands:
                if cmd == 'autocollect':
                    robot = Robot(server)
                    robot.collect(server.read_world().position)
                else:
                    print server.perform_actions([cmd])[0]

### \[roboclient.py\]

    from sys import argv
    from copy import deepcopy
    from time import sleep
    from Pyro.core import getProxyForURI
    
    class RoutingError(Exception):
        pass
    
    class Robot:
        def __init__(self, server):
            self.server = server
            
        def findPath(self, world, target_position):
            map = deepcopy(world.map)
            map[world.position[0]][world.position[1]] = [world.position]
            
            def propagate(pos):
                deltas = [(0, 1), (0, -1), (1, 0), (-1, 0)]
                new_positions = []
                current_path = map[pos[0]][pos[1]]
                
                for d in deltas:
                    new_pos = (pos[0] + d[0], pos[1] + d[1])
                    new_path = list(current_path) + [new_pos]
                    
                    try:
                        location = map[new_pos[0]][new_pos[1]]
                    except IndexError:
                        location = '#'
                    
                    if location != '#':
                        if location == ' ' or (len(location) > len(new_path)):
                            map[new_pos[0]][new_pos[1]] = new_path
                            new_positions.append(new_pos)
                return list(set(new_positions))
            
            positions = [world.position]
                
            while True:
                new_positions = []
                for pos in positions:
                    new_positions += propagate(pos)
                positions = new_positions
                
                if not positions:
                    raise RoutingError()
                
                path = map[target_position[0]][target_position[1]]
                if isinstance(path, list):
                    break
            return path
        
        def moveTo(self, position):
            world = self.server.read_world()
            path = self.findPath(world, position)[1:]
            return [self.stepTo(pos) for pos in path]
        
        def stepTo(self, pos):
            world = self.server.read_world()
            
            dr = pos[0] - world.position[0]
            dc = pos[1] - world.position[1]
            
            directions = {
                (0, -1): 'west',
                (0, 1): 'east',
                (1, 0): 'south',
                (-1, 0): 'north'}
            
            action = directions.get((dr, dc), None)
            if action is None:
                raise RoutingError()
            return self.performAction(action)
        
        def performAction(self, action):
            print self.server.read_world()
            print 'Robot action:', action
            sleep(1)
            self.server.perform_actions([action])
        
        def collect(self, position):
            while True:
                world = self.server.read_world()
                items = [i for i in world.items if i != position]
                
                if not items:
                    break
                
                self.moveTo(items[0])
                self.performAction('pickup')
                self.moveTo(position)
                self.performAction('drop')
            print world
        
        def play(self):
            self.collect(self.server.read_world().entry)
    
    if __name__ == '__main__':
        server = getProxyForURI(
            'PYROLOC://%s:%s/adventure_server' % (argv[1], argv[2]))
        Robot(server).play()

[Category:LOR-contest](Category:LOR-contest "wikilink")
