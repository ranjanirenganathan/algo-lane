Check for balanced parentheses in linear time using constant space  Related To: Strings
Problem
Example:

((()())(())) should be accepted but ())() should be rejected.​
Answer
Initialize a counter at zero. Whenever you see (, increase it by one, and whenever you see ), decrease it by one. Accept if the counter was always nonnegative and ended up at zero. Since counter is always between −1 and n, only O(log n) bits (note it is not O(1)) are needed to store it. So I don't think we can do O(1) space.

<?php

  // php code is wrapped in <?php tags

  class validate{
    
    public function __constructor(){
      
    }
    
    public function isValid(string $inputString){
      //initialize a counter
      $count = 0;
      $result = null;
      for($i=0;$i<strlen($inputString);$i++){
        if($inputString[$i] == '('){
          $count++;
        }else{
          $count--;
        }
      }
      if(empty($count)){
        $result = "Accepeted";
      }
      else{
        $result = "Rejected";
      }
      
      return $result;
      
  }
  }
  
  $test = new validate();
  $testString = '((()()()';
  echo "Checking if paranthesis are balanced :".$test->isValid($testString);
      

?>
