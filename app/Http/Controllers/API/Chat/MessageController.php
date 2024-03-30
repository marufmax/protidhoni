<?php

namespace App\Http\Controllers\API\Chat;

use App\Enums\MessageStatus;
use App\Events\AnythingEvent;
use App\Events\MessageSent;
use App\Models\Message;
use App\Models\User;
use Illuminate\Http\Request;

class MessageController
{
    public function send(User $receiver, Request $request)
    {
        if ($receiver->id === null) {
            return response()->json([
                'error' => 'Sending to invalid user',
            ]);
        }

//        $message = Message::create([
//            'contents' => $request->get('contents'),
//            'receiver' => $receiver,
//            'status' => MessageStatus::Sent,
//        ]);

        broadcast(new AnythingEvent($request->get('contents')));

        return response()->json([
            'message' => 'message sent',
        ]);
    }
}
