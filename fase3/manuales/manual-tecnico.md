# Traffic Splitting
## Cómo funcionan las métricas de oro, cómo puedes interpretar estas 5 pruebas de faulty traffic, usando como base los gráficos y métricas que muestra el tablero de Linkerd Grafana.

### Split 1 - Rabbit 100%
Segun las graficas generadas por Linkerd, se pudo observar que se manejaron las siguientes Golden Metrics
- RequesPerSecond: 0.3
- SuccessRate: 100%
- Latency: 18-20ms
![Split1](https://drive.google.com/uc?export=view&id=13pAWnrrSt-I1_voyAk0KmUHTiIqyyhRK)

### Split 2 - Kafka 100%
Segun las graficas generadas por Linkerd, se pudo observar que se manejaron las siguientes Golden Metrics
- RequesPerSecond: 0.3
- SuccessRate: 100%
- Latency: 30ms
![Split2](https://drive.google.com/uc?export=view&id=1bZy9RirEnw0QbinBTiPUsxHhPahjjykC)

### Split 3 - Rabbit 50% - Faulty 50%
Segun las graficas generadas por Linkerd, se pudo observar que se manejaron las siguientes Golden Metrics
- RequesPerSecond: 0.17
- SuccessRate: 60%
- Latency: 18-20ms
![Split3](https://drive.google.com/uc?export=view&id=1d6mvX2HuphTFd7oKQOfaRKksDLgRBMro)

### Split 4 - Kafka 50% - Faulty 50%
Segun las graficas generadas por Linkerd, se pudo observar que se manejaron las siguientes Golden Metrics
- RequesPerSecond: 0.03
- SuccessRate: 50%
- Latency: 30ms
![Split4](https://drive.google.com/uc?export=view&id=1UtHAKObSi8FOpd9MX1RfwEYS9_rHKi6Q)

### Split 5 - Rabbit 50% - Kafka 50%
Segun las graficas generadas por Linkerd, se pudo observar que se manejaron las siguientes Golden Metrics
#### Rabbit
- RequesPerSecond: 0.3
- SuccessRate: 100%
- Latency: 10ms
#### Kafka
- RequesPerSecond: 0.3
- SuccessRate: 100%
- Latency: 30ms
![Split5](https://drive.google.com/uc?export=view&id=1bHIPSGUg4hXifnv-PRKo1xSyVi9Bc87M)
![Split5_Rabbit](https://drive.google.com/uc?export=view&id=1gsiwJ9wjzn8fYW6Yq64O393Yz0DI8YGm)
![Split5_Kafka](https://drive.google.com/uc?export=view&id=1S5oK1__jU3lHoW8jI7ANwt1UWSrI7XVt)

## Menciona al menos 3 patrones de comportamiento que hayas descubierto.
- Los request por que fueron recibidos por Kafka tuvieron un mayor tiempo de respuesta en comparacion con RabbiMQ
- Cuando se trabajo con faulty traffic, RabbitMQ obtuvo un mayor porcentaje de request, en comparacion con Kafka que estuvo balanceado con un 50%
- Los request se fueron atendidos de mejor manera cuando el traffic split estuvo en un 50% para cada servicio de mensajeria, Kafka y RabbitMQ.

# RPC AND BROKERS
## ¿Qué sistema de mensajería es más rápido?
Ambos sistemas de mensajería cumplen con los requisitos básicos por lo que para dar una respuesta sobre cual es mejor tenemos que centrarnos en los detalles de cada mensajería y por el performance que tiene cada una en los puntos en común.
Hablando un poco sobre los puntos específicos, RabbitMQ tiene un mejor desempeño cuando se habla sobre encolar datos. Por otro lado, Apache Kafka soporta streaming de eventos mientras que RabbitMQ no tiene dicho soporte.
Por otro lado, el ruteo que proporciona RabbitMQ es más completo en relación con el ruteo de mensajes que ofrece __Apache Kafka__.
Tomando en cuenta estos parámetros y otros mostrados en el estudio realizado por Confluent, se puede concluir que Apache Kafka tiene un mejor desempeño en comparación a __RabbitMQ__.

## ¿Cuántos recursos utiliza cada sistema? (Basándose en los resultados que muestra el Dashboard de Linkerd)


## ¿Cuáles son las ventajas y desventajas de cada sistema?
### Ventajas y Desventajas de Apache Kafka y RabbitMQ
#### Apache Kafka
##### Ventajas
- Baja latencia
- Tolerancia a fallas
- Durabilidad de los mensajes
- Alta escalabilidad
##### Desventajas
- No tiene un set de monitoreo completo
- No tiene soporte al momento de ingresar mal el nombre del topic, únicamente obtiene los mensajes de un topic exacto
- Reduce el performance cuando se tienen varios brokers y consumers
#### RabbitMQ
##### Ventajas
- Permite dar distintas prioridades a los mensajes recibidos
- Escalabilidad
- Fácil uso
##### Desventajas
- No provee ordenamiento de mensajes
- No garantiza la atomicidad
- Los mensajes no persisten, una vez son transmitidos, son eliminados
### Ventajas y Desventajas de RPC y BROKERS
#### RPC
##### Ventajas
- Alta velocidad
- Posibilidad de realizar llamadas a terceros
##### Desventajas
- No tiene escalabilidad lineal
- El uso de threads y memoria se ve limitado cuando hacemos llamadas a terceros
#### BROKERS
##### Ventajas
- Maneja altos volúmenes de información
- Al estar basado en colas, almacena los mensajes o peticiones que llegan al servidor en el disco, por lo que tiene una alta tolerancia a un mayor número de peticiones
##### Desventajas
- No es tan rápido en comparación con RPC

## ¿Cuál es el mejor sistema?
Si bien el sistema de RPC nos proporciona una alta velocidad al momento de procesar las solicitudes, debemos de tomar en cuenta que, al momento de recibir una alta cantidad de peticiones, nuestro sistema puede colapsar ya que los recolectores de basura utilizan memoria y consumen recursos que hacen que nuestros procesos se detengan momentáneamente esto hace que nuestro sistema se sature y tengamos error del lado del cliente.
Este tipo de problema desaparece al momento de hacer uso de brokers. Estos brokers proporcionan una cola de procesos por lo que es muy poco probable que presentemos errores del lado del cliente, no obstante, esto tiene un coste de velocidad, por lo que podemos solucionar el problema de memoria y procesamiento, pero vamos a tener una disminución en la velocidad de procesamiento de peticiones. 
Estos problemas únicamente son observados al manejar grandes cantidades de datos, por lo que para aplicaciones de poco tráfico de peticiones, cualquier de los dos sistemas funcionara de la manera esperada.

# NOSQL DATABASES
## Redis
Redis es un proyecto de estructura de datos en memoria de código abierto que implementa una base de datos de clave-valor en memoria distribuida con durabilidad opcional.
Popular plataforma de datos en memoria utilizada como caché, intermediario de mensajes y base de datos que se puede implementar en las instalaciones, en las nubes y en entornos híbridos.

## TiDB
TiDB es una base de datos NewSQL distribuida de procesamiento transaccional/analítico híbrido (HTAP) de código abierto que admite las sintaxis MySQL y Spark SQL. Implementa la interfaz de redis.

## ¿Cuál de las dos bases (Redis y Tidis) se desempeña mejor y por qué?
La base de datos que se desempeña de una mejor forma es Redis, porque es una base de datos que se viene implementando hace tiempo y es un almacén de estructura de datos en memoria de código abierto (con licencia BSD), que se utiliza como base de datos, caché y agente de mensajes. Admite estructuras de datos como cadenas, hashes, listas, conjuntos, conjuntos ordenados con consultas de rango, mapas de bits, hiperloglogs, índices geoespaciales con consultas de radio y flujos. Redis tiene replicación integrada, secuencias de comandos Lua, desalojo de LRU, transacciones y diferentes niveles de persistencia en disco, y proporciona alta disponibilidad a través de Redis Sentinel y partición automática con Redis Cluster.

# CHAOS ENGINEERING
## ¿Cómo cada experimento se refleja en el gráfico de Linkerd, qué sucede?
### Trafico 50% - 50%
Cuando el trafico es generado con un 50% para cada ruta, tiende a mantener una igualdad en el trafico de peticiones, cabe mencionar que no se mantiene una exactitud
del 50% en cada ruta pero, el __LoadBalancer__ busca mantener la relacion 50/50 por lo que siempre se tendra un porcentaje aproximado al 50%.

### Trafico 100% - 0%
Cuando nosotros decidimos mantener una relacion de trabajo en la cual solo un servicio tendra la carga completa de todo el sistema, se puede observar que los servicios
se estresan ya que no soportan de manera adecuada la cantidad de peticiones que se estan realizando en ese momento. Mientras que el servicio sobre el cual se aplico
el 0% del trafico permanece en reposo con sus recursos intactos. Esto sucede al contrario de cuando optamos por la medida de trafico 50/50 ya que las peticiones
estan repartidas entre 2 rutas por lo que el trafico sera manejado por 2 paths esto hace que teoricamente se reduzca a la mitad el uso de recursos.

## ¿Qué diferencia tienen los experimentos?
Los experimentos estan diferenciados por el impacto que causan en el servicio, se explican a continuacion
- __Pod kill__: En este experimento un pod es eliminado lo que representa una carga en el sistema ya que este pod debe de ser creado nuevamente junto con todos los
contenedores que este Pod posea. A pesar de matar el pod, este experimento vuelve a levantar el pod de manera casi instantanea por lo que el daño provocado
al sistema no es tan grande.
- __Pod failure__: Este experimento funciona de manera distinta al __Pod kill__ ya que proboca un fallo en el pod, pero este no es eliminado sino que solo 
reinicia su activdad, a pesar de tener que reiniciar el pod, este tipo de experimento es mas dañino ya que inhabilita el uso del pod.
- __Container kill__: Este experimento funciona de manera similar al __Pod kill__ ya que ambos se encargan de destruir un elemento dentro del sistema, a pesar de ser
similares en comportamiento, __Container kill__ tiene un impacto mucho mayor ya que, unicamente destruye un contenedor dentro de un pod, mientras que __Pod kill__
destruye todos los contenedores que contiene dicho pod, haciendo que si este pod tiene una gran cantidad de contenedores, sea mas tardia su recuperacion.
- __Network Emulation (Netem) Chaos__: Este experimento no afecta la estructura del sistema en si, pero si produce un aumento en la latencia de respuesta de las
peticiones realizadas.
- __DNS Chaos__: El pod tendra una inyeccion de errores de tipo dns.

## ¿Cuál de todos los experimentos es el más dañino?
De los 5 experimentos realizados, PodFailure y DNSChaos parecieron ser los más dañinos, a diferencia de Pod Kill el cual reiniciaba el Pod instantáneamente,
Pod Failure e DNS Chaos inhabilitan los servicios en su totalidad, uno arrojando error por falla de Pod y el otro por erro de DNS. Por otra parte Slow Network a
pesar de afectar la latencia, permite que el sistema siga funcionando correctamente
