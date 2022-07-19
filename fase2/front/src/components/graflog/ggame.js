import {React, useState} from 'react';
import { ReactChartJs } from '@cubetiq/react-chart-js';

function Games ({logs = [] })  {

  var valores = [0,0,0,0,0]
    logs.map((uso) => {
      if (uso.game_id == '1'){
        valores[0]++;
      }else if(uso.game_id == '2'){
        valores[1]++;
      }else if(uso.game_id == '3'){
        valores[2]++;
      }else if(uso.game_id == '4'){
        valores[3]++;
      }else if(uso.game_id == '5'){
        valores[4]++;
      }else{}
    } );

    const data = {
        labels: ['Game1','Game2', 'Game3', 'Game4', 'Game5'],
        datasets: [{
          
          data: valores,
          backgroundColor: [
            'red',
            'blue',
            'yellow',
            'green',
            'orange'
            
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

export default Games;