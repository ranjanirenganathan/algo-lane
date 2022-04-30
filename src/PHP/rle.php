/**
Run Length Encoding (RLE) Data Compression Algorithm

Run–length encoding (RLE) is a simple form of lossless data compression that runs on sequences with the same value occurring many consecutive times. 
For below example
WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW
With a run–length encoding (RLE) data compression algorithm applied to the above hypothetical scan line, 
it can be rendered as 12W1B12W3B24W1B14W. This can be interpreted as a sequence of twelve W’s, one B, twelve W’s, three B’s, etc.
**/





<?php

  // php code is wrapped in <?php tags
function printRLE($str)
{
  $n = strlen($str);
  for($i =0;$i <$n;$i++){
    $count =1;
    while($i <$n-1 && $str[$i] == $str[$i+1]){
      $count++;
      $i++;
    }
    echo($str[$i].$count);
  }
   
}
 
// Driver Code
$str =  "aaaAAAAdebbBBccCC";
printRLE($str);
?>
