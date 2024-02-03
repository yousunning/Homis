package calc;

 public class Calc2 {
  public static void main(String[] args){
    System.out.println(2<<10);  // 2 << 11 10000000000
     
    System.out.println(11<<18); 
    // 3을 이진수로 나타내면 11
    // 이유 :
    // 2^1 자리: 1 (2^1 = 2)
    // 2^0 자리: 1 (2^0 = 1) 11 // 틀림

    double a = 3;
    double b = 18;

    double result = Math.pow(a, b);

     // 소수점 이하 1자리까지 출력
     Math.pow(3,18);
     System.out.printf("%.1f\n", result);
     
    //  base / exponent
    // System.out.println(result);

  }
 }
