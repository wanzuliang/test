/*




*/
import java.net.*;
import java.io.*;


public class MyServerTest extends Thread{
    private ServerSocket serverSocket;

    public MyServerTest(int port) throws IOException{
        serverSocket = new ServerSocket(port);
    }

    public void run(){
        while(true){
            try{
                System.out.println("Link Waiting...");
                Socket server = serverSocket.accept();

                DataInputStream in = new DataInputStream(server.getInputStream());
                System.out.println(in.readUTF());

                DataOutputStream out = new DataOutputStream(server.getOutputStream());

                out.writeUTF("Seccess : \n "
                     + "from " + server.getLocalSocketAddress()
                     + "to " + server.getRemoteSocketAddress()
                     + "\nBye.....#66ccff");

                out.flush();
                out.close();
                server.close();
            }catch(IOException e){
                e.printStackTrace();
                break;
            }
        }
    }

    public static void main(String[] args) {
        int port = 6666;
        try
        {
            Thread t = new MyServerTest(port);
            t.run();
        }catch(IOException e)
        {
            e.printStackTrace();
        }
    }
}