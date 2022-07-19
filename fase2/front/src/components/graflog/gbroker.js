import {React, useState} from 'react';
import { ReactChartJs } from '@cubetiq/react-chart-js';

function Broker ({logs = [] })  {
    var valores = [0,0]
    logs.map((uso) => {
      if (uso.queue == 'RabbitMQ' || uso.queue == 'rabbitMQ'){
        valores[0]++;
      }else if(uso.queue == 'Kafka' || uso.queue == 'kafka'){
        valores[1]++;
      }
    } );

    const data = {
        labels: ['RabbitMQ','Kafka'],
        datasets: [{
          
          data: valores,
          backgroundColor: [
            'red',
            'blue'
            
          ],
          borderColor: [
            'rgb(255, 99, 132)',
            'rgb(255, 159, 64)'
          ],
          borderWidth: 1
        }]
      };

      const config = {
        type: 'bar',
        data: data,
        options: {
          scales: {
            y: {
              beginAtZero: true
            }
          }
        },
      };

      return (
        <div>
          <ReactChartJs
              chartConfig={config}
          />
        </div>
      );

}

export default Broker;