require 'sinatra'
require 'json'

get '/saludo' do
  nombre = params['nombre']
  if nombre
    content_type :json
    { mensaje: "Hola, #{nombre}" }.to_json
  else
    status 400
    content_type :json
    { error: 'El parámetro "nombre" es requerido' }.to_json
  end
end