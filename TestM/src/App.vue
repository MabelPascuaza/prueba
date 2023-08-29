<template>
 
   <div class="centered-container">
    <h3 style="font-size: 3rem; color: rgb(60, 44, 131);">Consulta de datos</h3>
    
    <div class="content">
      <p>
        Para subir los datos haz click en el botón y espera la descarga en tu consola<br>
      </p>
    <button class="clickable-button" @click="executeMyGoFunction">Subir datos</button>
   </div>

    <p class="contact-message" style="font-size: 1rem;">
            Puedes contactarme por correo electrónico
      <a href="mailto:mabel.pascuaza@correounivalle.edu.co">aquí</a><br>
          </p>
    <p class="style1">Made by Mabel Pascuaza</p>
  </div>  
  
</template>



<script>

import axios from "axios";

export default {
  methods: {
    executeMyGoFunction() {
      axios.post("http://localhost:3000/execute-go-function")
        .then(response => {
          console.log(response.data); 
        })
        .catch(error => {
          console.error(error);
        });
    },
  },

};

</script>



<style scoped>

h3 {
  font-size: 3rem;
  text-align: center;
  margin-bottom: 150px;;
  
}

.centered-container{
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  text-align: center;
  box-sizing: border-box;
}

.content {
  text-align: center;
  margin: 20px 0;
}

.clickable-button {
  margin: 10px auto;
  padding: 10px 20px;
  background-color: rgb(59, 138, 160);
  color: white;
  cursor: pointer;
  border:none;
  border-radius: 4px;
  box-shadow: 2px 2px 4px rgba(0, 0, 0, 0.2); 
  transition: background-color 0.3s, color 0.3s; 
}

.clickable-button:hover {
  background-color: rgb(130, 140, 227);
}

.clickable-button:active {
  background-color: rgb(150, 206, 203);
}

.contact-message {
  font-size: 2rem;
  margin-top: 120px;
}

.style1{
  font-size: 0.5rem;
  color:rgb(115, 32, 153);
}

p {
  font-size: 1rem;
  text-align: center;
}

a {
    color:rgb(153, 46, 32);
}

</style>