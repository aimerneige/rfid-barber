using System;
using System.Windows;
using Microsoft.Extensions.DependencyInjection;
using RfidBarberClient.Services;

namespace RfidBarberClient;
/// <summary>
/// Interaction logic for App.xaml
/// </summary>
public partial class App : Application
{
    /// <summary>
    /// Gets the current <see cref="App"/> instance in use
    /// </summary>
    public new static App Current => (App)Application.Current;

    /// <summary>
    /// Gets the <see cref="IServiceProvider"/> instance to resolve application services.
    /// </summary>
    public IServiceProvider Services { get; }

    public App()
    {
        Services = ConfigureServices();

        InitializeComponent();
    }

    /// <summary>
    /// Configures the services for the application.
    /// </summary>
    private static IServiceProvider ConfigureServices()
    {
        var services = new ServiceCollection();

        services.AddHttpClient("RfidBarberClient.Http", client =>
        {
            client.BaseAddress
                = new("http://localhost:");
            client.DefaultRequestHeaders.Authorization
                = new("ACCESS_TOKEN");

        });
        services.AddSingleton<ICardService, CardService>();

        return services.BuildServiceProvider();
    }


}
