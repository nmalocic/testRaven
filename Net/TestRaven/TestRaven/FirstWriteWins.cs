using Raven.Client.Documents;
using Raven.Client.Exceptions;

namespace TestRaven;

public static class FirstWriteWins
{
    public static void FirstWrite(IDocumentStore store)
    {
        var item1 = new Item()
        {
            ID = "1",
            Value = "original"
        };

        var item2 = new Item()
        {
            ID = "1",
            Value = "updated"
        };
        
        using (var session = store.OpenSession())
        {
            session.Store(item1);
            session.SaveChanges();
        }

        try
        {
            using (var session = store.OpenSession())
            {
                session.Store(item2, changeVector: "", id: "1");
                session.SaveChanges();
            }
        }
        catch (ConcurrencyException _)
        {
            Console.WriteLine("Document was not saved, because it would override existing document");
        }
    }
}