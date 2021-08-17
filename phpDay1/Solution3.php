<?php

function convertStringToCamel($string){
	$result="";
	for($index=0;$index<strlen($string);$index++){
		if($string[$index]=='_'){
			continue;
		}
		if($string[$index-1]=='_'){
			$result.=strtoupper($string[$index]);
		}else{
			$result.=strtolower($string[$index]);
		}
	}
	return $result;
}
function snakeToCamel($array){
	for($index=0;$index<count($array);$index++){
		$array[$index]=convertStringToCamel($array[$index]);
	}
	return $array;
}

$input= ["snake_case", "camel_case"];
$output=snakeToCamel($input);
echo implode(",",$output);
