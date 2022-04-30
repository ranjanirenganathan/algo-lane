<?php

  // php code is wrapped in <?php tags

  class mergeSort{
    protected array $result = [];
    public function __constructor(){
      
    }
    
    public function merge(array $inputA, array $inputB){
      $i =$j=0;
      while($i<count($inputA) && $j <count($inputB)){
        if($inputA[$i]<= $inputB[$j] && !isset($this->result[$i])){
          $this->result[$i] =$inputA[$i];
          $i++;
          
        }else{
          $this->result[$j] = $inputB[$j];
          $j++;
        }
      }
      while($i <count($inputA) && !isset($this->result[$i])){
        $this->result[$i] = $inputA[$i];
        $i++;
        $k++;
      }
      while($j <count($inputB) && !isset($this->result[$j])){
        $this->result[$j] = $inputB[$j];
        $j++;
       
      }
      
      return $this->result;
      
    }
  }

  $inputA =[10,15,22,80];
  $inputB = [5,8,11,15,70,90];
  $test = new mergeSort();
echo (print_r($test->merge($inputA, $inputB),true));
?>
