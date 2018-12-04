// GreetingClient.java
 
import java.net.*;
import java.io.*;
 
public class MySocketTest
{
   public static void main(String [] args)
   {
      String serverName = "182.61.17.114";
      int port = 6666;
      try
      {
         System.out.println("Link Start: " + serverName + " , Port : " + port);
         Socket client = new Socket(serverName, port);
         System.out.println("Address : " + client.getRemoteSocketAddress());

         OutputStream outToServer = client.getOutputStream();
         DataOutputStream out = new DataOutputStream(outToServer);
         out.writeUTF("Link from " + client.getLocalSocketAddress());

         InputStream inFromServer = client.getInputStream();
         DataInputStream in = new DataInputStream(inFromServer);
         System.out.println("get: " + in.readUTF());

         in.close();
         out.flush();
         out.close();
         client.close();
      }catch(IOException e)
      {
         e.printStackTrace();
      }
   }
}