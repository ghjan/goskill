import java.util.stream.IntStream;  
public class compute_test{
        public static void main(String args[]){
        int sum = IntStream
             .range(0, 1000).parallel().map(n -> n * n).sum();
        System.out.println(sum);
    }
}