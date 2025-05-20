using Microsoft.AspNetCore.Mvc;

namespace SaludoApi.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class SaludoController : ControllerBase
    {
        [HttpGet("saludo")]
        public IActionResult Saludo([FromQuery] string nombre)
        {
            if (!string.IsNullOrEmpty(nombre))
            {
                return Ok(new { mensaje = $"Hola, {nombre}" });
            }
            else
            {
                return BadRequest(new { error = "El parámetro \"nombre\" es requerido" });
            }
        }
    }
}