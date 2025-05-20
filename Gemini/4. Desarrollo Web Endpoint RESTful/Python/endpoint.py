from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/saludo', methods=['GET'])
def saludo():
    nombre = request.args.get('nombre')
    if nombre:
        return jsonify({'mensaje': f'Hola, {nombre}'})
    else:
        return jsonify({'error': 'El parámetro "nombre" es requerido'}), 400

if __name__ == '__main__':
    app.run(debug=True)