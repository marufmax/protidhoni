<?php

namespace App\Http\Controllers\API;

use App\Http\Controllers\Controller;
use App\Http\Requests\LoginRequest;
use App\Models\User;
use Illuminate\Http\Response;
use Illuminate\Support\Facades\Auth;

class LoginController extends Controller
{
    private const TOKEN_PREFIX = 'api_token';
    public function create(LoginRequest $request)
    {
        if (!Auth::once($request->only(['email', 'password']))) {
            return response()->json([
                'error' => 'invalid user',
            ], Response::HTTP_UNAUTHORIZED);
        }

        $user = User::where(['email' => $request->get('email')])->firstOrFail();

        return response()->json([
            'id' => $user->id,
            'email' => $user->email,
            'token' => $user
                ->createToken(sprintf('%s_of_%s', self::TOKEN_PREFIX, $user->email))
                ->plainTextToken,
        ], Response::HTTP_CREATED);
    }
}
