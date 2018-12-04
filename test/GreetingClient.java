// GreetingClient.java
 
import java.net.*;
import java.io.*;
 
public class GreetingClient
{
   public static void main(String [] args)
   {
      String serverName = "localhost";
      int port = 6060;
      try
      {
         System.out.println("Link to : " + serverName + " , Port : " + port);
         Socket client = new Socket(serverName, port);
         System.out.println("Address : " + client.getRemoteSocketAddress());

         OutputStream outToServer = client.getOutputStream();
         DataOutputStream out = new DataOutputStream(outToServer);
 
         // out.writeUTF("Hello from " + client.getLocalSocketAddress());

         InputStream inFromServer = client.getInputStream();
         DataInputStream in = new DataInputStream(inFromServer);

         // System.out.println("get : " + in.readUTF());
         in.close();
         out.close();
         client.close();
      }catch(IOException e)
      {
         e.printStackTrace();
      }
   }
}