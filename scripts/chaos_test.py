import time
import subprocess

def kill_node(name):
    print(f"Derrubando nó: {name}")

if __name__ == "__main__":
    kill_node("Node-02")
    time.sleep(5)
    print("Verificando se um novo líder foi eleito...")
