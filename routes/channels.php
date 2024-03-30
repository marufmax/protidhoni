<?php

use App\Models\User;
use Illuminate\Support\Facades\Broadcast;

Broadcast::channel('App.Models.User.{id}', function ($user, $id) {
    dd($user);
    return (int) $user->id === (int) $id;
});

Broadcast::channel('private.chat.{id}', function (User $user, $id) {
   return true;
});

Broadcast::channel('anything', static function () {
    return true;
});
