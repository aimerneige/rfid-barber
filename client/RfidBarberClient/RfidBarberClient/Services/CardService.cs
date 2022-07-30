using RfidBarberClient.Models;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;
using Microsoft.Extensions.Logging;

namespace RfidBarberClient.Services;

public class CardService : ICardService
{
    private ILogger<CardService> _logger;
    private IHttpClientFactory _httpClientFactory;
    public CardService(ILogger<CardService> logger, 
        IHttpClientFactory httpClientFactory)
        => (_logger, _httpClientFactory) = (logger, httpClientFactory);

    public Task AddCardAsync()
    {
        throw new NotImplementedException();
    }

    public Task<IList<RfidCard>> GetAllCardAsync()
    {
        throw new NotImplementedException();
    }

    public Task<RfidCard> GetCardAsync(string id)
    {
        throw new NotImplementedException();
    }

    public Task RemoveCardAsync(RfidCard card)
    {
        throw new NotImplementedException();
    }

    public Task<RfidCard> UpdateCardAsync(RfidCard card)
    {
        throw new NotImplementedException();
    }
}
