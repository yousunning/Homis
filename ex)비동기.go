public class AsyncDigimon{
    public static void main(String[] args){
        for(int i=0;i<5;i++){
            Thread thread = new Thread(new PrintTask(i));
            thread.start();
        }
    }

    private static class PrintTask implements Runnable {
        final int taskNumber;
        ArrayList<String> digimon = new ArrayList<String>();
        
        public PintTask(int taskNumber){
            this.taskNumber = taskNumber;
            digimon.Add("안")
            digimon.Add("녕")
            digimon.Add("디")
            digimon.Add("지")
            digimon.Add("몬")
        }
    }

    @Override
    public void run() {
        System.out.println(digimon.get(taskNumber));
    }
}
// 몬
// 안
// 지
// 디
// 녕
