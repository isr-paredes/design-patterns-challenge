
from abc import ABC, abstractmethod

class BuildingFlyweight(ABC):
    @abstractmethod
    def render(self, unique_properties):
        pass

class ConcreteBuildingFlyweight(BuildingFlyweight):
    def __init__(self, shared_properties):
        self._shared_properties = shared_properties

    def render(self, unique_properties):
        # Combine shared properties with unique properties to render the building
        combined_properties = {**self._shared_properties, **unique_properties}
        return f"Rendering building with properties: {combined_properties}"
class BuildingFlyweightFactory:
    def __init__(self):
        self._flyweights = {}

    def get_building(self, shared_properties):
        # Use the shared properties as a key to find or create a flyweight object
        key = tuple(sorted(shared_properties.items()))
        if key not in self._flyweights:
            self._flyweights[key] = ConcreteBuildingFlyweight(shared_properties)
        return self._flyweights[key]

def main():
    # Create a flyweight factory
    factory = BuildingFlyweightFactory()

    # Define shared properties for buildings
    shared_properties = {
        "texture": "brick",
        "material": "concrete",
    }

    # Create and render buildings with unique properties
    for i in range(1000):
        unique_properties = {
            "height": f"{i} meters",
            "color": "blue" if i % 2 == 0 else "red",
        }
        building = factory.get_building(shared_properties)
        print(building.render(unique_properties))

if __name__ == "__main__":
    main()
