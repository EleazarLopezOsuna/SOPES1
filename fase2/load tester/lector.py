from random import random, randrange
from sys import getsizeof
import json

# Esta clase nos ayudara a manejar todas las acciones de lectura de los datos del archivo
class Reader():

    # Constructor de la clase
    def __init__(self):
        # En esta variable almacenaremos nuestros datos
        self.array = []
        
    # Obtener un valor random del array
    # NOTA: ESTO QUITA EL VALOR DEL ARRAY.
    def getData(self):
        # Obtenemos el numero de elementos del array
        length = len(self.array)
        
        # Si aun hay valores en el array
        if (length > 0):
            return self.array.pop(0)

        # Si ya no hay mas datos que leer del archivo
        else:
            # Imprimimos que ya no hay datos en el archivo
            print (">> Reader-Locust: No hay más valores para leer en el archivo.")
            # Retornamos nada
            return None
    
    # Cargar el archivo de datos json
    def load(self):
        # Mostramos en consola que estamos a punto de cargar los datos
        print (">> Reader-Locust: Iniciando con la carga de datos")
        # Ya que leeremos un archivo, es mejor realizar este proceso con un Try Except
        try:
            path = input("Reader-Locust: Ingrese la ruta del archivo JSON a cargar: ")
            with open(path, 'r') as dataFile:
                self.array = json.loads(dataFile.read())
            print (f'>> Reader-Locust: Datos cargados correctamente, {len(self.array)} datos -> {getsizeof(self.array)} bytes.')
        except Exception as e:
            print (f'>> Reader-Locust: No se cargaron los datos {e}')