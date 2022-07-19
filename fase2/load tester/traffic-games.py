from random import randrange
from locust import HttpUser, task, between
from lector import Reader
import json

# Clase principal que manejara locust
class Traffic(HttpUser):
    # Tiempo de espera entre peticiones
    # En este caso, esperara un tiempo de 0.1 segundos a 0.9 segundos entre cada llamada HTTP
    wait_time = between(0.1, 0.9)

    def on_start(self):
        self.reader = Reader()  # Instancia de la clase que nos permetira leer el archivo de carga
        self.dataSend = {}  # Variable que almacenara el json a enviar
        self.reader.load()  # Realizamos la lectura del archivo de carga
        self.rutas = ["/play","/match/addMatch"]
        self.endpoint = "" # Endpoint al que realizara la petición

    @task
    def initTraffic(self):
        # Obtener uno de los valores que enviaremos
        self.dataSend = self.reader.getData()

        # Si nuestro lector de datos nos devuelve None, es momento de parar
        if (self.dataSend is not None):
            data_to_send = json.dumps(self.dataSend)
            self.endpoint = self.rutas[randrange(0,2)]
            response = self.client.post(self.endpoint, json=self.dataSend)
            respuesta = json.loads(response.content)
            print(">> Reader-Locust-Response: ", respuesta)
        else:
            print(">> Reader-Locust: Envio de tráfico finalizado")
            self.stop(True) # Se envía True como parámetro para el valor "force", este fuerza a locust a parar el proceso inmediatamente.
