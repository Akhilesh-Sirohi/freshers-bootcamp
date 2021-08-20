<?php

namespace App\Http\Controllers;

use App\User;
use Illuminate\Http\Request;
use Illuminate\Http\Response;
use Illuminate\Support\Facades\DB;

class UsersController extends Controller
{

    public function store(Request $request)
    {
        $request->validate([
            'first_name'=>'required',
            'last_name'=>'required',
            'email'=>'required'
        ]);

        $user=array(
            'first_name'=>$request->get('first_name'),
            "last_name"=>$request->get('last_name'),
            "email"=>  $request->get('email'),
        );
        $id = DB::table('users')->insertGetId($user);
        return response()->json([
            'id' => $id ,
            'first_name'=>$user['first_name'],
            'last_name'=>$user['last_name'],
            'email'=>$user['email'],
        ]);
    }

    public function getUserById($id)
    {
        $users=DB::table('users')
            ->where('id',$id)->get();
        if(count($users)==0){
            return response()->json([
                'message' => "user with id-$id does not exist"
            ]);
        }
        return response()->json([
            $users[0]
        ]);
    }
    public function getAllUsers()
    {
        return response()->json([
            'All users' => DB::table('users')->get(),
        ]);
    }

    public function deleteUserById($id)
    {
        DB::table('users')
            ->where('id' , $id)
            ->delete();

        return response()->json([
            'message' => "user with id-$id deleted"
        ]);
    }
}
