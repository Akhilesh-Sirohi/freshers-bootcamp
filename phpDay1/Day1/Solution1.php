<?php

function array_flatten($array){
    $flattened_array=[];
    foreach($array as $element){
        if (is_array($element)){
            $flattened_array=array_merge($flattened_array,array_flatten($element));
        } else{
            array_push($flattened_array,$element);
        }
    }
    return $flattened_array;
}

$array=[1,2,3,[4,5],6,7,[8,9,[11,12],13],14];
$output=array_flatten($array);
echo implode(",",$output);