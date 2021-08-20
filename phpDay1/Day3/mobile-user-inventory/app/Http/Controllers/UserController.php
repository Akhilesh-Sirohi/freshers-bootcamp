<?php

namespace App\Http\Controllers;

use App\User;
use http\Params;
use Illuminate\Http\Request;
use Illuminate\Http\Response;
use Illuminate\Support\Facades\DB;
use Illuminate\Database\QueryException;
use Nette\Schema\ValidationException;
use Illuminate\Support\Facades\Validator;
use App\Services\UserService;


class UserController extends Controller
{
    public function hello(Request $request)
    {
        // return $id;
        $query = $request->query;
        //dd($query->keys());
        $key = $query->keys();
        //return $key;
        return $query->get($key[0]);
        //dd($request);
        return "hello";
        //return $name;
    }

    public function createUser(Request $request): \Illuminate\Http\JsonResponse
    {
        $data = $request->all();
        $userObj = new UserService();
        return $userObj->createUser($data);
    }


    public function getUserByParameter(Request $request)
    {
        $userObj = new UserService();
        $query = $request->query;
        $key = $query->keys();
        //dd($key);
        if (count($key) == 0) {
            return $userObj->getAllUsers();
        }
        return $userObj->getUserByParameter($key[0], $query->get($key[0]));
    }

    public function deleteUserByParameter(Request $request)
    {
        $userObj = new UserService();
        $query = $request->query;
        $key = $query->keys();
        //dd($key);
        if (count($key) == 0) {
            return response()->json("bad request", 400);
        }
        return $userObj->deleteUserByParameter($key[0], $query->get($key[0]));
    }
}
