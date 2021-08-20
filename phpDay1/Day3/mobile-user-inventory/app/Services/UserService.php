<?php

namespace App\Services;

use Illuminate\Http\Request;
use Illuminate\Http\Response;
use Illuminate\Support\Facades\DB;
use Illuminate\Database\QueryException;
use Nette\Schema\ValidationException;
use Illuminate\Support\Facades\Validator;
use Psy\Util\Json;


class UserService
{

    public function createUser($data)
    {
        $validator = Validator::make($data, [
            'name' => 'required',
            'email' => 'required|email',
            'mobile_number' => 'required|regex:/^[6-9][0-9]{9}$/'
        ]);
        if ($validator->fails()) {
            return response()->json($validator->errors(), 400);
        }


        try {
            $id = DB::table('users')->insertGetId($data);
        } catch (QueryException $e) {
            return response()->json($e->errorInfo, Response::HTTP_BAD_REQUEST);
        }
        return response()->json([
            "status" => 200,
            "message" => "User Created!",
            "data" => [
                'id' => $id,
                'name' => $data['name'],
                'mobile_number' => $data['mobile_number'],
                'email' => $data['email'],
            ]
        ]);
    }

    public function getAllUsers()
    {
        return response()->json(
        //'All users' => DB::table('users')->get(),
            DB::table('users')->get()
        );
    }

    public function getUserByParameter($parameter, $value)
    {
        $users = DB::table('users')
            ->where($parameter, $value)->get();
        if (count($users) == 0) {
            return response()->json([
                'message' => "user with $parameter -$value does not exist"
            ]);
        }
        return response()->json([
            $users
        ]);
    }

    public function deleteUserByParameter($parameter, $value)
    {
        DB::table('users')
            ->where($parameter, $value)
            ->delete();

        return response()->json([
            "status" => 200,
            'message' => "User with $parameter-$value deleted"
        ]);
    }
}
