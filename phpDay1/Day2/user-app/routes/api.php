<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\UsersController;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:api')->get('/users', function (Request $request) {
    return $request->user();
});

Route::post('/users',  [UsersController::class, 'store']); //....... add user
Route::post('/hello',function (){
    return "hii";
});
Route::delete('/users/{id}',[UsersController::class, 'deleteUserById']);
Route::get('/users/{id}',[UsersController::class, 'getUserById']);
Route::get('/users',[UsersController::class, 'getAllUsers']);


