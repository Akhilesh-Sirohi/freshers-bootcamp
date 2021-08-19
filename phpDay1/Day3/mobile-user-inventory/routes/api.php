<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\UserController;

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

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});

Route::post('users',[UserController::class,'createUser']);

Route::get('users',[UserController::class,'getAllUsers']);
Route::get('users/{id}',[UserController::class,'getUserById']);
Route::get('users/name/{name}',[UserController::class,'getUserByName']);
Route::get('users/mobile_number/{mobile_number}',[UserController::class,'getUserByMobileNumber']);
Route::get('users/email/{email}',[UserController::class,'getUserByEmail']);

Route::delete('users/{id}',[UserController::class,'deleteUserById']);
Route::delete('users/name/{name}',[UserController::class,'deleteUserByName']);
Route::delete('users/mobile_number/{mobile_number}',[UserController::class,'deleteUserByMobileNumber']);
Route::delete('users/email/{email}',[UserController::class,'deleteUserByEmail']);
