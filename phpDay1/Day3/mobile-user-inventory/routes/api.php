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

Route::post('users', [UserController::class, 'createUser']);

//Route::get('users', [UserController::class, 'getAllUsers']);

Route::get('users/{id?}/{name?}/{email?}/{mobile_number?}',[UserController::class,'getUserByParameter']);

Route::delete('users/{id?}/{name?}/{email?}/{mobile_number?}',[UserController::class,'deleteUserByParameter']);
