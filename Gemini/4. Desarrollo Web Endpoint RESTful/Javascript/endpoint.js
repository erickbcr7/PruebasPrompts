const express = require('express');
const app = express();
const port = 3000;

app.get('/saludo', (req, res) => {
  const nombre = req.query.nombre;
  if (nombre) {
    res.json({ mensaje: `Hola, ${nombre}` });
  } else {
    res.status(400).json({ error: 'El parámetro "nombre" es requerido' });
  }
});

app.listen(port, () => {
  console.log(`Servidor escuchando en http://localhost:${port}`);
});