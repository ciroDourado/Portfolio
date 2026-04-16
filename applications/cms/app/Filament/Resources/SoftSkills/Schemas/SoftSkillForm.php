<?php

namespace App\Filament\Resources\SoftSkills\Schemas;

use Filament\Forms\Components\Textarea;
use Filament\Forms\Components\TextInput;
use Filament\Schemas\Components\Section;
use Filament\Schemas\Schema;

class SoftSkillForm
{
    public static function configure(Schema $schema): Schema
    {
        return $schema
            ->components([
                Section::make('Skill Details')
                    ->schema([
                        TextInput::make('title')
                            ->required()
                            ->maxLength(255)
                            ->label('Title'),
                        TextInput::make('icon')
                            ->required()
                            ->maxLength(255)
                            ->label('Icon'),
                        TextInput::make('order')
                            ->numeric()
                            ->default(0)
                            ->label('Order'),
                        Textarea::make('description')
                            ->required()
                            ->columnSpanFull()
                            ->label('Description'),
                    ])->columns(3),
            ]);
    }
}
