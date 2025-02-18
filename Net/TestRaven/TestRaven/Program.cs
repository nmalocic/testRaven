// See https://aka.ms/new-console-template for more information

using Raven.Client.Documents;
using Raven.Client.Exceptions;
using TestRaven;

Console.WriteLine("Hello, World!");
const string dbName =  "testdapr";


static IDocumentStore GertStore()
{
    return (new DocumentStore
    {
        Urls = new[] { "http://localhost:8080" },
        Database = dbName
    }).Initialize();
}

var store = GertStore();

Transaction.CheckTransaction(store);
//FirstWriteWins.FirstWrite(store);


