using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using CommunityToolkit.Mvvm.ComponentModel;

namespace RfidBarberClient.Models;
internal class MainWindowModel : ObservableObject
{
    private uint _cardCount;
    private uint _weekIncome;
    private double _weekIncomeTrend;

    public uint CardCount
    {
        get => _cardCount;
        set => SetProperty(ref _cardCount, value);
    }

    public uint WeekIncome
    {
        get => _weekIncome;
        set => SetProperty(ref _weekIncome, value);
    }

    public double WeekIncomeTrend
    {
        get => _weekIncomeTrend;
        set => SetProperty(ref _weekIncomeTrend, value);
    }
}
