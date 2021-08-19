<?php

function getMaskedNumber($number){
	$masked_number=substr($number,0,2);
	$masked_number.=str_repeat("*",strlen($number)-4);
	$masked_number.=substr($number,-2,2);
	return $masked_number;
}

$number="9812345678";
$output=getMaskedNumber($number);
echo $output;