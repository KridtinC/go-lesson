public class Girl extends Person implements Singer {

    public String sing() {
        return "Lalala";
    }


    public Girl(String name, int age) {
        super(name, age);
    }
    

    @Override
    public String getName() {
        return "girl: " + this.name;
    }

   
}

