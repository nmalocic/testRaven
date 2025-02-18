using Raven.Client.Documents;

namespace TestRaven;

public static class Transaction
{
  public static void CheckTransaction(IDocumentStore store)
  {
    var item1 = new Item()
    {
      ID = "item1",
      Value = "notImportant"
    };

    var item2 = new Item()
    {
      ID = "item2",
      Value = "not importa"
    };
    try
    {
      using var session = store.OpenSession();
      session.Store(item1);
      session.Store(item2);
      session.Delete("item2");
      session.SaveChanges();
    }
    catch (Exception ex)
    {
      Console.WriteLine(ex.Message);
    }
}  
}