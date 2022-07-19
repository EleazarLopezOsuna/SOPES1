import React from 'react';
import '../App.css';
import {useState} from "react";

function Tree ({ data = [] }) {
    return (
        <>
       
          {data.map((tree) => (
                <TreeNodo node={tree} />
          ))}
        
      </>
    );
  };

const TreeNodo = ({node}) => {
    const [hijoVisible, sethijoVisible] = useState(false);
    const hasChild = node.hijo ? true : false;

    return (
        <>
        <tbody>
        <tr onClick={e => sethijoVisible((v) => !v)}>
                    <th> {node.pid}</th>   
                    <td>{node.nombre}</td>     
                    <td>{node.estado}</td>
        </tr>
        </tbody>
        {hasChild && hijoVisible && (
              <Tree data={node.hijo} />
        )}
        
        </>
    );
}

export default Tree;