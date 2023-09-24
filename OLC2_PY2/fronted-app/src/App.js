
import React, { useState } from 'react';
import './App.css';
import { MyButtons , Pestaña} from './Components/TextArea';
import { okaidia } from '@uiw/codemirror-theme-okaidia';
import CodeMirror from '@uiw/react-codemirror';
import { javascript } from '@codemirror/lang-javascript';


function App() {
  const [code, setCode] = useState('');
  const [selectedFile, setSelectedFile] = useState(null);

  const [items, setItems] = useState({ListSymbol: [], ListError: [] , ListConsole: []});

  const handleChange = (value) => {
    setCode(value);
  };
  // ----------------------------------------------------------------
 
  const handleFileChange = (event) => {
    const file = event.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = (e) => {
        setCode(e.target.result);
      };
      reader.readAsText(file);
    }
    setSelectedFile(file);
  };
 
  const Ejecutar = async () => {          
              fetch('http://localhost:5000/File', {
                    method: 'POST',
                    body: code,     
              })
              .then(response => response.json())
              .then(data => {
                console.log('Success:', data);
                    setItems(data);
              })
              .catch(error => {
                console.error('Error:', error);
              })
   
  };
  
  const [mostrarSVG, setMostrarSVG] = useState(false);

  const mostrarOcultarSVG = () => {
    setMostrarSVG(!mostrarSVG);
  }

  const GenerarAST = () => {
      fetch('http://localhost:5000/AST', {
        method: 'POST',   
      })
      .then((response) => response.text())
      .then((data) => {
               setMostrarSVG(data);
               console.log(data.ListConsole)
      })
      .catch((error) => {
        console.error('Error al obtener el AST:', error);
      });
  };

  return (
    <div className="App">
      <div className="Botones">
          <div className="file-input-wrapper">
              <input className ="file-input" type="file" id="fileInput" onChange={handleFileChange}/>
              <label className ="file-input-label" htmlFor="fileInput">  ABRIR  </label>
          </div>
          <MyButtons
              texto="EDITAR"
              manejarClic={() => console.log(code)}
          />
          <MyButtons
              texto="ACERCA DE"
              manejarClic={() => console.log("acerca de")}
          />
      </div>
      <div className="codemirror-style">
         <CodeMirror
          value={code}
          height="600px"
          width='900px'
          theme={okaidia}
          extensions={[javascript({ jsx: true })]}
          onChange={handleChange}
          />
      </div>
      <div className="Valores_Salida">
        
        <Pestaña items={items.ListError} env={items.ListSymbol} terminal={items.ListConsole}/>
      </div> 
      <div className="EjecutarReP">
        <MyButtons
              texto="EJECUTAR"
              manejarClic={Ejecutar}
          />
        <MyButtons
              texto="REPORTE"
              manejarClic={GenerarAST}
          />
           {mostrarSVG &&  (
              <MyButtons
                  texto="CERRAR ARBOL"
                  manejarClic={mostrarOcultarSVG}
              />)}
            
      </div>
      {mostrarSVG &&  ( 
        <div class='scroll-container'>
            <div className='svg'
              dangerouslySetInnerHTML={{ __html: mostrarSVG }}
            >
            </div>
        </div>
      ) }
    </div>
  );
}

export default App;
