

class Inventory:
    items: dict = {}

    def __init__(self):
        self.items = []

    def add_item(self, id, item):
        self.items[id] = item

    def remove_item(self, id):
        return self.items.pop(id)

    def has_item(self, item):
        return item in self.items
