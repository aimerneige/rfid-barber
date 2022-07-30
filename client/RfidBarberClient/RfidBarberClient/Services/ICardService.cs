using System.Collections.Generic;
using System.Threading.Tasks;
using RfidBarberClient.Models;

namespace RfidBarberClient.Services;

public interface ICardService
{
    public Task<IList<RfidCard>> GetAllCardAsync();
    public Task<RfidCard> GetCardAsync(string id);
    public Task AddCardAsync();
    public Task<RfidCard> UpdateCardAsync(RfidCard card);
    public Task RemoveCardAsync(RfidCard card);
}
