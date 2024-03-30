<?php

namespace App\Http\Controllers\API\Auth;

class PingController
{
    public function __invoke()
    {
        return [
          'ping' => 'pong',
        ];
    }
}
