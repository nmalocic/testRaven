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

var item1 = new Item()
{
    Id = "items/1",
    Value = "original"
};

var item2 = new Item()
{
    Id = "items/1",
    Value = "updated"
};

var store = GertStore();

using (var session = store.OpenSession())
{
    session.Store(item1);
    session.SaveChanges();
}

try
{
    using (var session = store.OpenSession())
    {
        session.Store(item2, changeVector: "", id: "items/1");
        session.SaveChanges();
    }
}
catch (ConcurrencyException _)
{
    Console.WriteLine("Document was not saved, because it would override existing document");
}

try
{
    using (var session = store.OpenSession())
    {
        session.Delete("items/1", "randomChangeVectorShouldFail");
        session.SaveChanges();
    }
}
catch (Exception e)
{
    Console.WriteLine("document not delted wrong change vector");
}


