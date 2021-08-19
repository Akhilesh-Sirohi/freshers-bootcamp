<?php
function GetValuesByParameter($json,$parameter){
    $values=[];
    foreach($json as $element){
        array_push($values,$element->{$parameter});
    }
    return $values;

}


function GetPlayerNameByAge($json,$age){
    $values=[];
    foreach($json as $element){
        if($element->{'age'}==$age){
            array_push($values,$element->{'name'});
        }
    }
    return $values;
}


$data="{\"players\":[{\"name\":\"Ganguly\",\"age\":45,\"address\":{\"city\":\"Hyderabad\"}},
    {\"name\":\"Dravid\",\"age\":45,\"address\":{\"city\":\"Hyderabad\"}},
    {\"name\":\"Dhoni\",\"age\":37,\"address\":{\"city\":\"Hyderabad\"}},
    {\"name\":\"Virat\",\"age\":35,\"address\":{\"city\":\"Hyderabad\"}},
    {\"name\":\"Jadeja\",\"age\":35,\"address\":{\"city\":\"Hyderabad\"}},
    {\"name\":\"Jadeja\",\"age\":35,\"address\":{\"city\":\"Hyderabad\"}}]}";
$json_players=json_decode($data)->{'players'};
//var_dump($json->{"players"});
$names=GetValuesByParameter($json_players,'name');
$age=GetValuesByParameter($json_players,'age');
$address=GetValuesByParameter($json_players,'address');
$city=GetValuesByParameter($address,'city');

print_r($names);
print_r($age);
print_r($city);
$unique_names=array_unique($names);
print_r($unique_names);

$max_age=max($age);
$max_age_players= GetPlayerNameByAge($json_players,$max_age);
print_r($max_age_players);