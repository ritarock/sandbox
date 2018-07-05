import sys
import json

def main(arg):
    f = open(arg[0], 'r')
    tmp = json.load(f)
    # [{'x1': '1', 'y1': '5'}, {'x2': '2', 'y2': '1o'}, {'x3': '3', 'y3': '15'}]

if __name__ == '__main__':
    main(sys.argv[1:])
