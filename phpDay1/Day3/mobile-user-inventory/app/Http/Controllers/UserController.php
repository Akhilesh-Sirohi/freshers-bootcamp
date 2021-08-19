<?php

namespace App\Http\Controllers;

use App\User;
use Illuminate\Http\Request;
use Illuminate\Http\Response;
use Illuminate\Support\Facades\DB;
use Illuminate\Database\QueryException;
use Nette\Schema\ValidationException;
use Illuminate\Support\Facades\Validator;


class UserController extends Controller
{

    public function createUser(Request $request)
    {
        $validator = Validator::make($request->all(), [
            'name' => 'required',
            'email' => 'required|email',
            'mobile_number'=>'required|regex:/^[6-9][0-9]{9}$/'
        ]);
        if ($validator->fails()) {
            return response()->json($validator->errors(),Response::HTTP_BAD_REQUEST);
        }

        $user=array(
            'name'=>$request->get('name'),
            "mobile_number"=>$request->get('mobile_number'),
            "email"=>  $request->get('email'),
        );
        try{
            $id = DB::table('users')->insertGetId($user);
        }
        catch (QueryException $e){
            return response()->json($e->errorInfo, Response::HTTP_BAD_REQUEST);

        }
        return response()->json([
            'id' => $id ,
            'name'=>$user['name'],
            'mobile_number'=>$user['mobile_number'],
            'email'=>$user['email'],
        ]);
    }
    public function getAllUsers()
    {
        return response()->json([
            'All users' => DB::table('users')->get(),
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

    public function getUserByEmail($email){
        $users= DB::table('users')->where('email',$email)->get();
        if(count($users)==0){
            return response()->json([
                'message' => "user with email -$email does not exist"
            ]);
        }
        return response()->json([
            $users[0]
        ]);
    }
    public function getUserByName($name){
        $users= DB::table('users')->where('name',$name)->get();
        if(count($users)==0){
            return response()->json([
                'message' => "user with name -$name does not exist"
            ]);
        }
        return response()->json([
            $users
        ]);
    }

    public function getUserByMobileNumber($mobile_number){
        $users= DB::table('users')->where('mobile_number',$mobile_number)->get();
        if(count($users)==0){
            return response()->json([
                'message' => "user with mobile number -$mobile_number does not exist"
            ]);
        }
        return response()->json([
            $users[0]
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
    public function deleteUserByName($name)
    {
        DB::table('users')
            ->where('name' , $name)
            ->delete();

        return response()->json([
            'message' => "user with name -$name deleted"
        ]);
    }
    public function deleteUserByMobileNumber($mobile_number)
    {
        DB::table('users')
            ->where('mobile_number' , $mobile_number)
            ->delete();

        return response()->json([
            'message' => "user with mobile number-$mobile_number deleted"
        ]);
    }
    public function deleteUserByEmail($email)
    {
        DB::table('users')
            ->where('email' , $email)
            ->delete();

        return response()->json([
            'message' => "user with email -$email deleted"
        ]);
    }
}
