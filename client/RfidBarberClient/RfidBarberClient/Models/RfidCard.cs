using System.Text.Json.Serialization;

namespace RfidBarberClient.Models;
#nullable disable
public class RfidCard
{
    [JsonPropertyName("rfid")]
    public string Id { get; set; }
    public string Name { get; set; }
    public string Phone { get; set; }
    public decimal Balance { get; set; }
}
#nullable restore
