import json
import random

MARK = ['Spade', 'Heart', 'Diamond', 'Club']
RANK = ['A', '2', '3', '4', '5', '6', '7', '8', '9', '10', 'J', 'Q', 'K']


class Deck:
    def __init__(self):
        self.deck = {}
        for m in MARK:
            self.deck[m] = RANK


    def draw(self):
        mark_no = MARK[random.randrange(4)]
        rank_no = RANK[random.randrange(len(self.deck[mark_no]))]
        self.deck[mark_no].remove(rank_no)
        return rank_no

class Player:
    def __init__(self, name):
        self.name = name
        self.hands = []

class Game:
    def __init__(self, player1, player2, deck):
        self.player1 = player1
        self.player2 = player2
        self.deck = deck
        for i in range(0, 2):
            self.player1.hands.append(self.deck.draw())
        for i in range(0, 2):
            self.player2.hands.append(self.deck.draw())


def main():
    pass
if __name__ == '__main__':
    main()
