package exam;

public class exam1 {

   public static void main(String[] args) {
      
      //2^11+3^18
      String two = "1";    //2진수값 받을 스트링
      String thr = "1";    //3진수값 받을 스트링
      int twoX = 2;    //밑 = 2
      int thrX = 3;    //밑 = 3
      int twoN = 11;    //지수 = 11
      int thrN = 18;    //지수 = 18
      int sum = 0;
      
      for (int i = 1; i<=twoN; i++) {    //2^11을 2진수로 표현
         two+="0";
      }
      System.out.println(two);
      
      for (int j = 1; j<=thrN; j++) {    //3^18을 3진수로 표현
         thr+="0";
      }
      System.out.println(thr);
      
      sum = Integer.parseInt(two,2) + Integer.parseInt(thr,3);
      
      System.out.println(Integer.parseInt(two,2)); //2진수 -> 10진수
      System.out.println(Integer.parseInt(thr,3)); //3진수 -> 10진수
      System.out.println(sum);
      
      


   }

}
