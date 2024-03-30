<?php

use App\Http\Controllers\API\Auth\LoginController;
use App\Http\Controllers\API\Auth\PingController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

Route::get('/user', static function (Request $request) {
    return $request->user();
})->middleware('auth:sanctum');

Route::get('ping', PingController::class);

## Auth
Route::post('login', [LoginController::class, 'login']);
Route::post('register', [LoginController::class, 'register']);

## Chat
Route::post('messages/{receiver}', [\App\Http\Controllers\API\Chat\MessageController::class, 'send']);

## Debug
Route::post('/debug', static function (Request $request) {
    $signature = "{$request->get('id')}:{$request->get('channel_name')}";

   return hash_hmac(
       'sha256',
       $signature,
       env('REVERB_APP_SECRET'),
   );
});
