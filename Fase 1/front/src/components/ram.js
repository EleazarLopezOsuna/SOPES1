import React from 'react';
import '../App.css';
import api_host from './conf';
import { ReactChartJs } from '@cubetiq/react-chart-js';

function Ram(){

    const [values, setValues] = React.useState(
        [
            {
                data : [
                    {virtualMachine : 0, total : 0, memoriaLibre: 0, memoriaEnUso : 0, porcentaje : 0},
                ],
            },
            {
                data : [
                    {virtualMachine : 100, total : 0, memoriaLibre: 0, memoriaEnUso : 0, porcentaje : 90},
                ],
            },
            {
                data : [
                    {virtualMachine : 0, total : 0, memoriaLibre: 0, memoriaEnUso : 0, porcentaje : 0},
                ],
            },
            {
                data : [
                    {virtualMachine : 100, total : 0, memoriaLibre: 0, memoriaEnUso : 0, porcentaje : 90},
                ],
            },
            {
                data : [
                    {virtualMachine : 50, total : 1, memoriaLibre: 1, memoriaEnUso : 1, porcentaje : 50}

                ],
            }
        ]
      );


      const [values2, setValues2] = React.useState(
        [
            {
                data : [
                    {virtualMachine : 0, total : 0, memoriaLibre: 0, memoriaEnUso : 0, porcentaje : 90},
                ],
            },
            {
                data : [
                    {virtualMachine : 100, total : 0, memoriaLibre: 0, memoriaEnUso : 0, porcentaje : 50},
                ],
            },
            {
                data : [
                    {virtualMachine : 0, total : 0, memoriaLibre: 0, memoriaEnUso : 0, porcentaje : 0},
                ],
            },
            {
                data : [
                    {virtualMachine : 100, total : 0, memoriaLibre: 0, memoriaEnUso : 0, porcentaje : 30},
                ],
            },
            {
                data : [
                    {virtualMachine : 50, total : 2, memoriaLibre: 2, memoriaEnUso : 2, porcentaje : 50}

                ],
            }
        ]
    );

      const [loadGraph, setLoadGraph] = React.useState(false);

      

    const getRAM = async (evt) => {
 
        //evt.preventDefault();
        var name = evt.target.name;
        var endpoint = "";

        if (name == 'r1'){
            endpoint = `${api_host.api_one}` + evt.target.value;
        } else { 
            endpoint = `${api_host.api_two}` + evt.target.value;
        }

        var lista = []
        while (lista.length < 5){
            const response =
            await fetch(endpoint, { headers: {'Content-Type': 'application/json'} } )
              var value = await response.json()
              
              if (name === 'r1' && value.mv === 'vm1'){
                
                lista.push(value)
            } else if(name === 'r2' && value.mv === 'vm2') { 
             
                lista.push(value)
            }
              
        }


        if (name == 'r1'){
                
            setValues(lista);
        } else { 
         
            setValues2(lista);
        }
      }
    
    return(
        <>
        <div class="box">
            <button class="button is-link" onClick={getRAM} name="r1" value="ram1">RAM VM1</button>
            <button class="button is-danger" onClick={getRAM} name="r2" value="ram2">RAM VM2</button>
            <div class="columns">
                
                <div class="column">
                    <p class="title is-2">Maquina Virtual </p>
                    <p class="title is-1">{values[values.length -1 ].mv}</p>
                    <p class="title is-1">{values2[values2.length -1 ].mv}</p>
                </div>
                <div class="column">
                    <p class="title is-2">Memoria Total </p>
                    <p class="title is-1">{values[values.length - 1].data[0].total}</p>
                    <p class="title is-1">{values2[values2.length - 1].data[0].total}</p>
                </div>
                <div class="column">
                    <p class="title is-2">Memoria Libre </p>
                    <p class="title is-1">{values[values.length - 1].data[0].memoriaLibre}</p>
                    <p class="title is-1">{values2[values2.length - 1].data[0].memoriaLibre}</p>
                </div>
                <div class="column">
                    <p class="title is-2">Memoria en Uso </p>
                    <p class="title is-1">{values[values.length -1 ].data[0].memoriaEnUso}</p>
                    <p class="title is-1">{values2[values2.length -1 ].data[0].memoriaEnUso}</p>
                    
                </div>
                <div class="column">
                    <p class="title is-2">Mem en Uso (%) </p>
                    <p class="title is-1">{values[values.length -1 ].data[0].porcentaje}</p>
                    <p class="title is-1">{values2[values2.length -1 ].data[0].porcentaje}</p>
                </div>
            
            </div>
        </div>
        <div class="box">
            <button class="button is-success" onClick={() => setLoadGraph(prevState => !prevState)}>
                SHOW GRAPH
            </button>
            
            {loadGraph ?<Graph data={{vm1: values, vm2 : values2}} /> : null}
           
        </div>
        </>
    );
    /*data2={values2.cuadro}*/
    

}

const Graph= ({data}) => {
    var gstate = data.vm1.map((uso) => {
        return parseInt(uso.data[0].porcentaje)
    } );
    var gstate2 = data.vm2.map((duso) => {
        return parseInt(duso.data[0].porcentaje)
    } );
    var config={
        type: 'line',
        options: {
            responsive: true,
            title: {
                display: true,
                text: 'Uso de RAM',
            },
            tooltips: {
                mode: 'index',
                intersect: false,
            },
            hover: {
                mode: 'nearest',
                intersect: true,
            },
        },
        data: {
            datasets: [
                {
                    label: 'vm1',
                    data: gstate,
                    fill: false,
                    borderColor: '#0000FF',
                },
                {
                    label: 'vm2',
                    data: gstate2,
                    fill: false,
                    borderColor: '#ff6384',
                }
            ],
            
            labels: ['1','2','3','4','5'],
            
        },
    }
    console.log(gstate)
    return (
      <div>
        <ReactChartJs
            chartConfig={config}
        />
      </div>
    );
}
export default Ram;