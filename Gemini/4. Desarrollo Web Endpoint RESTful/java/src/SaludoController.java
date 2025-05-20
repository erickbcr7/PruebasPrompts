import java.util.Map;

@RestController
public class SaludoController {

    @GetMapping("/saludo")
    public Map<String, String> saludo(@RequestParam(value = "nombre", required = false) String nombre) {
        if (nombre != null) {
            return Map.of("mensaje", "Hola, " + nombre);
        } else {
            return Map.of("error", "El parámetro \"nombre\" es requerido");
        }
    }
}